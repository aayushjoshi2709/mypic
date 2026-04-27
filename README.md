# mypic

My pic is a application to store all the momories of your loved once at a single place

# how to setup

- Create a env file in the server folder

    ```
    # Application configuration
    PORT=<backend port>

    # MongoDB configuration
    MONGO_DB_URL=<mongodb url>
    MONGO_DB_MAX_POOL_SIZE=<max pool size for mongdb>
    MONGO_DB_MIN_POOL_SIZE=<min pool size for mongdb>
    MONGO_DB_NAME=<dbname for mongodb to connect to>

    # Redis configuration
    REDIS_URI=<redis url>

    # AWS S3 configuration
    AWS_ACCESS_KEY_ID=<access key for aws>
    AWS_SECRET_KEY=<access key secret for aws>
    AWS_REGION=<access key secret for aws>
    AWS_S3_BUCKET_NAME=<bucket name for s3>
    AWS_CLOUD_FRONT_URL =<cloudfront url>

    # JWT configuration
    JWT_SECRET_KEY=<jwt secret for your application>
    JWT_EXPIRES_IN=<jwt expiration for your application>

    SITE_DOMAIN=<domain of your website, leave blank string if running locally>
    ```

- Create another env file in client folder
    ```
    VITE_BASE_URL=<frontend url>
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