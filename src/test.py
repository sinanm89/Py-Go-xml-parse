import requests
import datetime

def micro_to_ms(number):
    return number * 0.001

def get_it(uri):
    now = datetime.datetime.now()
    response = requests.get(uri)
    later = datetime.datetime.now()
    return micro_to_ms((later - now).microseconds)

def get_go():
    avg = 0
    for i in range(1,100):
        avg += get_it("http://go-vast-parse.mobworkz.com/65617")
    return avg/100

def get_py():
    avg = 0
    for i in range(1,100):
        avg += get_it("http://vast-parse.mobworkz.com/65617")
    return avg/100
