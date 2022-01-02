const redis = require('redis');

class RedisClient {
  constructor() {
    this.Client = null;
  }

  async connect() {
    this.Client = redis.createClient({url: 'redis://default:redisgrupo11@34.125.118.239:6379'});
    await this.Client.connect();
    return this.Client;
  }
}

module.exports.RedisClient = RedisClient;