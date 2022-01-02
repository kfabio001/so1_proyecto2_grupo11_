const connection = require('./redis-connection');

async function launch() {
  let conn = new connection.RedisClient();
  await conn.connect();
  console.log('launched...');
}

module.exports.launch = launch;