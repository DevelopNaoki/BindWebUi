from flask import Flask, url_for, redirect, render_template, request, session, Blueprint

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
            from ..module import SearchUser
            find = SearchUser(request.form["username"], request.form["password"])
            if find:
                session["user"] = format(str(request.form["username"]))

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

