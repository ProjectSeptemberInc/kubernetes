var _ = require('underscore');
_.mixin(require('underscore.string').exports());

exports.ipv4 = function (ocets, prefix) {
  return {
    ocets: ocets,
    prefix: prefix,
    toString: function () {
      return [ocets.join('.'), prefix].join('/');
    }
  }
};

exports.hostname = function hostname (n, prefix) {
  return _.template("<%= pre %>-<%= seq %>-<%= suf %>")({
    pre: prefix || 'core',
    suf: exports.get_suffix(),
    seq: _.pad(n, 2, '0'),
  });
};

exports.rand_string = function () {
  var crypto = require('crypto');
  var shasum = crypto.createHash('sha256');
  shasum.update(crypto.randomBytes(256));
  return shasum.digest('hex');
};


exports.rand_suffix = exports.rand_string().substring(50);

exports.get_suffix = function() {
  if (process.env['AZ_SUFFIX']) {
    return process.env['AZ_SUFFIX'];
  } else {
    return exports.rand_suffix;
  }
}

exports.join_output_file_path = function(prefix, suffix) {
  output_dir = process.env['AZ_OUTPUT_DIR'] || './output';
  return output_dir + '/' + [prefix, exports.get_suffix(), suffix].join('_');
};
