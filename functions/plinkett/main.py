from random import choice

quotes = []

with open("./plinkett.txt", "r") as fin:
    quotes = fin.readlines()

def plinkett(inp):
    return choice(quotes)

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": plinkett("")},
        ],
    }]
