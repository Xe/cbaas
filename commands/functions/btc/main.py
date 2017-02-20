import requests

def priceCheck():
    "Shows BTC exchange rates on Bitstamp"

    try:
        info = requests.get("https://www.bitstamp.net/api/ticker/").json()

        return "Bitstamp prices: Average: %s, High: %s, Low: %s" %\
                (info["ask"], info["high"], info["low"])
    except Exception as e:
        return "There was some error looking bitcoin prices: %s" % e.message

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": priceCheck()},
        ],
    }]
