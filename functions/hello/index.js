console.log("starting function");

exports.handle = function(e, ctx, cb) {
    console.log("processing event: %j", e);
    cb(null, [{
        "to": e.target,
        "body": [
            {"body": "hello "},
            {"mention": e.sender},
            {"body": "!"},
        ]
    }]);
};
