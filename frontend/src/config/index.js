var index = ''

if (process.env.MY_ENV) {
  index = require('./_index_' + process.env.MY_ENV)
} else {
  index = require('./_index_' + process.env.NODE_ENV)
}

export default index.default
