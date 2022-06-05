SELECT 'CREATE DATABASE hokkori'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'hokkori')\gexec
