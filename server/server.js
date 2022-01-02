const express = require("express");
const bodyParser = require("body-parser");
const MongoClient = require("mongodb").MongoClient;
const cors = require("cors");
const launchRedis = require("./redis/launcher");
const connection = require('./redis/redis-connection');
const app = express();

const connString = "mongodb://sopes1:managerl@34.125.118.239:27017";

app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

MongoClient.connect(connString, {
  useUnifiedTopology: true,
})
  .then(async (client) => {
    console.log("Conectado a BD!");
    const db = client.db("sopes1-data");
    const dataCollection = db.collection("registros");
    await launchRedis.launch();

    app.get("/", (req, res) => {
      db.collection("registros")
        .find()
        .toArray()
        .then((results) => {
          res.json(results);
        })
        .catch((error) => console.error(error));
    });

    app.get("/one_dose", (req, res) => {
      db.collection("registros")
        .find({ n_dose: 1 })
        .toArray()
        .then((results) => {
          res.json(results);
        })
        .catch((error) => console.error(error));
    });

    app.get("/two_dose", (req, res) => {
      db.collection("registros")
        .find({ n_dose: 2 })
        .toArray()
        .then((results) => {
          res.json(results);
        })
        .catch((error) => console.error(error));
    });

    app.get("/ninos", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const ninos = await conn.Client.get('ninos');
          //console.log(ninos);
          res.status(200).send(ninos);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });

      app.get("/adolescentes", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const adolescentes = await conn.Client.get('adolescentes');
          //console.log(adolescentes);
          res.status(200).send(adolescentes);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });

      app.get("/jovenes", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const jovenes = await conn.Client.get('jovenes');
          //console.log(jovenes);
          res.status(200).send(jovenes);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });

      app.get("/adultos", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const adultos = await conn.Client.get('adultos');
          //console.log(adultos);
          res.status(200).send(adultos);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });

      app.get("/vejez", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const vejez = await conn.Client.get('vejez');
          //console.log(vejez);
          res.status(200).send(vejez);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });

      app.get("/nombres", async (req, res) => {
        try {
          let conn = new connection.RedisClient();
          await conn.connect();
          const vejez = await conn.Client.lRange('lNombres', 0, -1);
          //console.log(vejez);
          res.status(200).send(vejez);
        } catch (err) {
            console.log(err);
            res.status(404).json({error: err})
        }
      });
  })
  .catch(console.error);

app.listen(3003, () => {
  console.log("Puerto 3003...");
});
