import path from 'path';
import { fileURLToPath } from 'url';
import { read } from 'node-yaml';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const configPath = path.join(__dirname, '../../configs', 'local.yaml');

const loadConfig = async () => {
  try {
    const config = await read(configPath);
    return config;
  } catch (error) {
    console.error(`Error reading or parsing YAML file: ${error}`);
    throw error;
  }
};

export default loadConfig;
