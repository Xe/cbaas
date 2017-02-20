from random import choice
import json

quotes = []

with open("./bobross.json", "r") as fin:
    print fin
    quotes = json.load(fin)

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": choice(quotes)},
        ],
    }]
