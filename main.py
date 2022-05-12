from flask import Flask

app = Flask(__name__)

from router import *

@app.route('/', methods=["GET"])
def index():
    return 'Hello main'

app.register_blueprint(api)

if __name__ == '__main__':
    from init import serverConfig
    serverConfig = serverConfig()

    if 'debug' in serverConfig and (serverConfig['debug'].upper() == "TRUE"):
      debug = True
    else:
      debug = False
    app.debug = debug

    app.run(host=serverConfig['listem'],port=serverConfig['port'])
