#!/bin/bash
redis-cli <<EOF
SET config:1 '{"packs_config":[250, 500, 1000, 2000, 5000]}'
SADD config 1
EOF

