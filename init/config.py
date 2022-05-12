import json

json_open = open('server.conf', 'r')
json_load = json.load(json_open)

def serverConfig():
  return json_load['server']
