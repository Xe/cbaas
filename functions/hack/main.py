import hack

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": hack.hack("")},
        ],
    }]
