CREATE TABLE IF NOT EXISTS weather_info (
    location VARCHAR(300) NOT NULL,
    day TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    temperature DECIMAL(10, 2) NOT NULL,
    humidity DECIMAL(10, 2) NOT NULL,
    wind_speed DECIMAL(10, 2),
    condition VARCHAR(300)
)

