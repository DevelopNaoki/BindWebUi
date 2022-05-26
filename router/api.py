from flask import Flask, url_for, redirect, render_template, request, session
from flask import Blueprint
from datetime import timedelta

app = Blueprint('app', __name__)

#app.secret_key = 'thisiscertificateauthority'
#app.permanent_session_lifetime = timedelta(minutes=3)

@app.route('/')
def redirectDashboard():
    if "user" in session:
        return redirect('/dashboard')
    else:
        return redirect('/login')

@app.route('/login')
def login():
    if "user" in session:
        return redirect('/dashboard')
    else:
        return render_template('login.html', title="Login")

@app.route('/dashboard')
def dashboard():
    if "user" in session:
        return render_template('dashboard.html', title="Certificate Authority User Interface", session=session)
    else:
        return redirect('/login')


@app.route('/api/login', methods=["POST"])
def loginApi():
    if "user" in session:
        return redirect('/dashboard')
    else:
        if "username" in request.form and "password" in request.form:
            import MySQLdb
            connection = MySQLdb.connect(
                host='192.168.0.251',
                user='root',
                passwd='root',
                db='certificateAuthority')
            cursor = connection.cursor()

            sql = ('''SELECT id FROM users WHERE userName=%s AND pass=%s''')
            param = (str(request.form["username"]), str(request.form["password"]))
            cursor.execute(sql, param)

            if len(cursor.fetchall()) == 1:
                session["user"] = format(str(request.form["username"]))

            connection.commit()
            connection.close()

            return redirect('/')
        else:
            return redirect('/login')

@app.route('/api/logout', methods=["GET"])
def logoutApi():
    if "user" in session:
        session.pop("user", None)
        return redirect('/login')
    else:
        return redirect('/login')

