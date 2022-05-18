from flask import render_template
from flask import Blueprint

api = Blueprint('api', __name__)

@api.route('/dashboard')
def app_b1():
    return render_template('dashboard.html')
    #return "hello B b1"
