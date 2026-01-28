-- PC Parts Tracker Database Schema

-- Create database (run manually)
-- CREATE DATABASE pcparts;

-- Parts table
CREATE TABLE IF NOT EXISTS parts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- GPU, CPU, RAM, Motherboard, Storage, PSU, Case
    brand VARCHAR(100) NOT NULL,
    specs TEXT,
    image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_parts_type ON parts(type);
CREATE INDEX idx_parts_name ON parts(name);

-- Dealers table
CREATE TABLE IF NOT EXISTS dealers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    url VARCHAR(500) NOT NULL,
    authenticity_rating DECIMAL(3,2) DEFAULT 0.0 CHECK (authenticity_rating >= 0 AND authenticity_rating <= 5),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Prices table
CREATE TABLE IF NOT EXISTS prices (
    id SERIAL PRIMARY KEY,
    part_id INTEGER NOT NULL REFERENCES parts(id) ON DELETE CASCADE,
    dealer_id INTEGER NOT NULL REFERENCES dealers(id) ON DELETE CASCADE,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    currency VARCHAR(3) DEFAULT 'USD',
    in_stock BOOLEAN DEFAULT true,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(part_id, dealer_id)
);

CREATE INDEX idx_prices_part_id ON prices(part_id);
CREATE INDEX idx_prices_dealer_id ON prices(dealer_id);
CREATE INDEX idx_prices_price ON prices(price);

-- Sample data for testing

-- Insert sample dealers
INSERT INTO dealers (name, url, authenticity_rating, is_verified) VALUES
('TechStore Direct', 'https://techstore.example.com', 4.8, true),
('PC Components Plus', 'https://pcplus.example.com', 4.5, true),
('Budget Hardware', 'https://budgetparts.example.com', 3.9, false),
('Elite PC Parts', 'https://elitepc.example.com', 4.9, true),
('Digital Depot', 'https://digitaldepot.example.com', 4.2, true);

-- Insert sample parts
INSERT INTO parts (name, type, brand, specs, image_url) VALUES
-- GPUs
('RTX 4090', 'GPU', 'NVIDIA', '24GB GDDR6X, 16384 CUDA Cores, Boost Clock: 2.52 GHz', ''),
('RTX 4080', 'GPU', 'NVIDIA', '16GB GDDR6X, 9728 CUDA Cores, Boost Clock: 2.51 GHz', ''),
('RX 7900 XTX', 'GPU', 'AMD', '24GB GDDR6, 96 CUs, Game Clock: 2.3 GHz', ''),
('RTX 4070 Ti', 'GPU', 'NVIDIA', '12GB GDDR6X, 7680 CUDA Cores, Boost Clock: 2.61 GHz', ''),

-- CPUs
('Ryzen 9 7950X', 'CPU', 'AMD', '16 Cores, 32 Threads, Base: 4.5 GHz, Boost: 5.7 GHz', ''),
('Core i9-14900K', 'CPU', 'Intel', '24 Cores, 32 Threads, Base: 3.2 GHz, Turbo: 6.0 GHz', ''),
('Ryzen 7 7800X3D', 'CPU', 'AMD', '8 Cores, 16 Threads, Base: 4.2 GHz, Boost: 5.0 GHz, 96MB Cache', ''),
('Core i7-14700K', 'CPU', 'Intel', '20 Cores, 28 Threads, Base: 3.4 GHz, Turbo: 5.6 GHz', ''),

-- RAM
('Corsair Vengeance DDR5 32GB', 'RAM', 'Corsair', '32GB (2x16GB), DDR5-6000, CL30', ''),
('G.Skill Trident Z5 32GB', 'RAM', 'G.Skill', '32GB (2x16GB), DDR5-6400, CL32', ''),
('Kingston Fury Beast 32GB', 'RAM', 'Kingston', '32GB (2x16GB), DDR5-5600, CL36', ''),

-- Motherboards
('ASUS ROG Strix X670E-E', 'Motherboard', 'ASUS', 'AMD X670E, Socket AM5, ATX, PCIe 5.0, DDR5', ''),
('MSI MAG Z790 Tomahawk', 'Motherboard', 'MSI', 'Intel Z790, LGA 1700, ATX, PCIe 5.0, DDR5', ''),
('Gigabyte B650 AORUS Elite', 'Motherboard', 'Gigabyte', 'AMD B650, Socket AM5, ATX, PCIe 4.0, DDR5', ''),

-- Storage
('Samsung 990 PRO 2TB', 'Storage', 'Samsung', '2TB NVMe SSD, PCIe 4.0, Read: 7450 MB/s, Write: 6900 MB/s', ''),
('WD Black SN850X 1TB', 'Storage', 'Western Digital', '1TB NVMe SSD, PCIe 4.0, Read: 7300 MB/s, Write: 6300 MB/s', ''),
('Crucial P3 Plus 4TB', 'Storage', 'Crucial', '4TB NVMe SSD, PCIe 4.0, Read: 5000 MB/s, Write: 4200 MB/s', '');

-- Insert sample prices
-- RTX 4090 prices
INSERT INTO prices (part_id, dealer_id, price, currency, in_stock) VALUES
(1, 1, 1599.99, 'USD', true),
(1, 2, 1649.99, 'USD', true),
(1, 3, 1549.99, 'USD', false),
(1, 4, 1699.99, 'USD', true),
(1, 5, 1589.99, 'USD', true);

-- RTX 4080 prices
INSERT INTO prices (part_id, dealer_id, price, currency, in_stock) VALUES
(2, 1, 1199.99, 'USD', true),
(2, 2, 1229.99, 'USD', true),
(2, 4, 1249.99, 'USD', false),
(2, 5, 1189.99, 'USD', true);

-- RX 7900 XTX prices
INSERT INTO prices (part_id, dealer_id, price, currency, in_stock) VALUES
(3, 1, 999.99, 'USD', true),
(3, 2, 1019.99, 'USD', true),
(3, 3, 979.99, 'USD', true),
(3, 5, 989.99, 'USD', false);

-- Ryzen 9 7950X prices
INSERT INTO prices (part_id, dealer_id, price, currency, in_stock) VALUES
(5, 1, 549.99, 'USD', true),
(5, 2, 569.99, 'USD', true),
(5, 3, 539.99, 'USD', false),
(5, 4, 579.99, 'USD', true);

-- Add more sample prices for other parts
INSERT INTO prices (part_id, dealer_id, price, currency, in_stock) VALUES
(6, 1, 589.99, 'USD', true),
(7, 1, 449.99, 'USD', true),
(8, 2, 419.99, 'USD', true),
(9, 1, 179.99, 'USD', true),
(10, 3, 189.99, 'USD', true),
(11, 4, 149.99, 'USD', true),
(12, 1, 329.99, 'USD', true),
(13, 2, 279.99, 'USD', true),
(14, 3, 189.99, 'USD', true),
(15, 1, 199.99, 'USD', true),
(16, 2, 129.99, 'USD', true),
(17, 4, 249.99, 'USD', true);