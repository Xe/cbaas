"use strict";

const randomItem = function (arr) {
	  if (!Array.isArray(arr)) {
		    throw new TypeError('Expected an array');
	  }

	  return arr[Math.floor(Math.random() * arr.length)];
};

const pf = require ("./facts.json");

exports.handle = function(ev, ctx, cb) {
    cb(null, [
        {
            to: ev.target,
            body: [
                {body: randomItem(pf)}
            ]
        }
    ]);
};
