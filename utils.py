from datetime import datetime
import socket
import os
import socket


def get_subject():
    current_time = datetime.now().strftime("%Y%m%d")
    return f"ip:{get_local_ip()}/{current_time}"

# def get_local_ip():
#     hostname = socket.gethostname()
#     ip = socket.gethostbyname(hostname)
#     return ip


def get_local_ip():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect(("8.8.8.8", 80))
    ip = s.getsockname()[0]
    return ip
