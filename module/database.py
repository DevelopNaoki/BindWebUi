import MySQLdb

def ConnectDB():
    connection = MySQLdb.connect(
      host='192.168.0.251',
      user='root',
      passwd='root',
      db='certificateAuthority')
    return connection

def  CloseDB(connection)
    connection.commit()
    connection.close()

def SearchUser(name, password):
    connection = ConnectDB()
    cursor = connection.cursor()

    sql = ('''SELECT id FROM users WHERE userName=%s AND pass=%s''')
    param = (str(username), str(password))
    cursor.execute(sql, param)

    CloseDB(connection)

    if len(cursor.fetchall()) == 1:
      find = true
    else:
      find = false

    return find
