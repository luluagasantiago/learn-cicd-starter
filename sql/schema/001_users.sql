-- +goose Up
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;




if (bono == 100 OR bono == 135 ){
    do this
}


OR 

1 1 -> 1 
1 0 -> 1 
0 1 -> 1  
0 0 -> 0