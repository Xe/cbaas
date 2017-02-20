import random

def handle(event, context):
    splitline = event["bodyString"].split()
    choices = " ".join(splitline[1:])
    choices = choices.split(", ")

    choice = random.choice(choices)

    return [{
        "to": event["to"],
        "body": [
            {"body": choice},
        ],
    }]
