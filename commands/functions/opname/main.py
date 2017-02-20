from opname import opname

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": opname("")},
        ],
    }]
