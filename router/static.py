from flask import render_template
from flask import Blueprint

static = Blueprint('api', __name__)

@static.route('/css')
def cssFile():
    #return render_template('COMP_B/index_B_1.html')
    return "This is CSS File"

@static.route('/js')
def jsFile():
    return "This is JavaScript File"
