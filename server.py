from flask import Flask
from router import *

webApp = Flask(__name__)

webApp.register_blueprint(app)

if __name__ == '__main__':
    from init import serverConfig
    serverConfig = serverConfig()

    if 'debug' in serverConfig and (serverConfig['debug'].upper() == "TRUE"):
      webApp.debug = True
    else:
      webApp.debug = False

    if 'ssl_certificate' in serverConfig and 'ssl_private_key' in serverConfig:
      import ssl
      sslContext = ssl.SSLContext(ssl.PROTOCOL_TLSv1_2)
      sslContext.load_cert_chain(
        serverConfig['ssl_certificate'], serverConfig['ssl_private_key']
      )
      webApp.run(host=serverConfig['listem'],port=serverConfig['port'],ssl_context=sslContext)
    else:
      webApp.run(host=serverConfig['listem'],port=serverConfig['port'])
