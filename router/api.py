from flask import render_template
from flask import Blueprint

api = Blueprint('api', __name__)

@api.route('/api')
def app_b1():
    #return render_template('COMP_B/index_B_1.html')
    return "hello B b1"
