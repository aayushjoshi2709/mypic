# mypic

My pic is a application to store all the momories of your loved once at a single place

# how to setup

- Create a env file in the server folder

    ```
    # Application configuration
    PORT=3000

    # MongoDB configuration
    MONGO_DB_URL=
    MONGO_DB_MAX_POOL_SIZE=
    MONGO_DB_MIN_POOL_SIZE=
    MONGO_DB_NAME=

    # Redis configuration
    REDIS_URI=

    # AWS S3 configuration
    AWS_ACCESS_KEY_ID=
    AWS_SECRET_KEY=
    AWS_REGION=
    AWS_S3_BUCKET_NAME=
    AWS_CLOUD_FRONT_URL =

    # JWT configuration
    JWT_SECRET_KEY=
    JWT_EXPIRES_IN=

    SITE_DOMAIN=""
    ```

- Create another env file in client folder
    ```
    VITE_BASE_URL=http://localhost:3000
    ```

- Install client dependencies
    ```
    cd client && npm install
    ```

- Install server dependencies
    ```
    cd server && go mod download
    ```


# Running dev env

- Run frontend server
    ```
    cd client && npm run dev    
    ```

- Run backend server
    ```
    cd server && make dev
    ```