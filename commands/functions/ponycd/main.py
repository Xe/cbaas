import pony

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": pony.when("")},
        ],
    }]
