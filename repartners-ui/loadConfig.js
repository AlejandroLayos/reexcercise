const fs = require('fs');
const path = require('path');

function loadConfig() {
  const env = process.env.NODE_ENV || 'development';
  const configPath = path.resolve(__dirname, 'config.json');
  const config = JSON.parse(fs.readFileSync(configPath, 'utf-8'));

  if (config[env]) {
    Object.keys(config[env]).forEach(key => {
      process.env[key] = config[env][key];
    });
  } else {
    console.error(`No configuration found for environment: ${env}`);
  }
}

module.exports = loadConfig;