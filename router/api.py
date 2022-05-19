from flask import Flask, url_for, redirect, render_template, request, session
from flask import Blueprint
from datetime import timedelta

app = Blueprint('app', __name__)

app.secret_key = 'abcdefghijklmn'
app.permanent_session_lifetime = timedelta(minutes=3)

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
        return render_template('dashboard.html', title="Bind WEB User Interface", session=session)
    else:
        return redirect('/login')


@app.route('/api/login', methods=["POST"])
def loginApi():
    if "user" in session:
        return redirect('/dashboard')
    else:
        return redirect('/login')
