import socket

host = socket.gethostname()
port = 25  # The same port as used by the server
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))
s.sendall(b'Helo test.example.com')
data = s.recv(1024)
s.close()
print('Received', repr(data))