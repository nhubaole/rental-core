import pkg from 'pg';
import loadConfig from "../config/config.js";

const { Pool } = pkg;

let pool;

const initializeDatabase = async () => {
  try {
    const config = await loadConfig();
    const connectionString = `postgres://${config.db.user}:${config.db.password}@${config.db.host}:${config.db.port}/${config.db.name}`;

    pool = new Pool({
      connectionString: connectionString,
    });

    pool.on('connect', () => {
      console.log('Database connection success');
    });
  } catch (error) {
    console.error('Failed to initialize the database:', error);
  }
};

await initializeDatabase();

export const query = (text, params) => {
  if (!pool) {
    throw new Error("Database pool has not been initialized yet.");
  }
  return pool.query(text, params);
};
