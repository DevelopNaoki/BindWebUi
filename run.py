from flask import Flask
import json
from routes import app
from datetime import timedelta

webApp = Flask(__name__)
webApp.register_blueprint(app)
webApp.secret_key = 'thisiscertificateauthority'
webApp.permanent_session_lifetime = timedelta(minutes=3)

def serverConfig():
  json_open = open('server.conf', 'r')
  json_load = json.load(json_open)
  return json_load['server']

if __name__ == '__main__':
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



