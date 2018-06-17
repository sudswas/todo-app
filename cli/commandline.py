import argparse
from ConfigParser import SafeConfigParser
import json
import requests
import sys

def main():
    parser = argparse.ArgumentParser(description='Add and retrieve to-dos')
    parser.add_argument('--get', metavar='get', dest='get_val', type=str, choices=['all','today'],
                        help='Get to-do lists by literals.')
    parser.add_argument('--add', metavar='add', dest='add_val', type=str, nargs='+',
                        help='Add a to-do as: --add <task> <date>')
    parser.add_argument('--getbydate', metavar='getdate', dest='get_by_date', type=str,
                        help='--getbydate <YYYY-MM-DD>')

    args = parser.parse_args(sys.argv[1:])
    handle = Handler()
    if args.add_val:
        if len(args.add_val) < 2:
            sys.exit("We need two values for an add operation in the format: --add <task> <date>")
        handle.handle_add(args.add_val)

    elif args.get_val:
        handle.handle_get(args.get_val)

    elif args.get_by_date:
        handle.handle_get_by_date(args.get_by_date)

class Handler(object):

    def __init__(self):
        config = SafeConfigParser()
        config.read('config.ini')
        self.url = config.get('default', 'url')

    def handle_get(self, value):
        if value == "all":
            # Call the server with gettodo
            response = requests.get(self.url + "/todo")
            if response.content != "null":
               print response.content
            else:
               print "There are no to-dos, please create some"
        elif value == "today":
            # Call server with gettodo/today
            response = requests.get(self.url + "/todo/today")
            if response.content:
               print response.content
            else:
               print "There are no to-dos for today"

    def handle_get_by_date(self, value):
        ## TODO: Add logic to validate the date format
        # Call server with todo/<date>
        response = requests.get(self.url + "/todo/" + value)
        print response.content

    def handle_add(self, values):
        ## TODO: Add logic to validate the date format
        # Call server with todo/<date>
        data = {}
        data['task'] = values[0]
        data['dueby'] = values[1]
        response = requests.post(self.url + "/todo", data=json.dumps(data))
        if response.status_code == 200:
           print "Request successfully sent"

if __name__ == "__main__":
    main()

