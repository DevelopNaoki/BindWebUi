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
      app.debug = True
    else:
      app.debug = False

    app.run(host=serverConfig['listem'],port=serverConfig['port'])
