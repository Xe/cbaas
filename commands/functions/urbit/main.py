import urbit

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": urbit.urbit(event["sender"] + event["bodyString"])},
        ],
    }]
