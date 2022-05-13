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

    if 'ssl_certificate' in serverConfig and 'ssl_private_key' in serverConfig:
      import ssl
      ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLSv1_2)
      ssl_context.load_cert_chain(
        serverConfig['ssl_certificate'], serverConfig['ssl_private_key']
      )
      app.run(host=serverConfig['listem'],port=serverConfig['port'],ssl_context=ssl_context)
    else:
      app.run(host=serverConfig['listem'],port=serverConfig['port'])
