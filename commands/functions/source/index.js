console.log("starting function");

exports.handle = function(e, ctx, cb) {
    console.log("processing event: %j", e);
    cb(null, [{
        "to": e.target,
        "body": [
            {"body": "source code for this bot: https://github.com/Xe/cbaas"},
        ]
    }]);
};
