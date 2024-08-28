import random

def generate_ip():
    return f"{random.randint(0,255)}.{random.randint(0,255)}.{random.randint(0,255)}.{random.randint(0,255)}"

def generate_large_ip_file(filename, size_gb):
    target_size = size_gb * 1024 * 1024 * 1024  # Convert GB to bytes
    current_size = 0

    with open(filename, 'w') as f:
        while current_size < target_size:
            ip = generate_ip() + '\n'
            f.write(ip)
            current_size += len(ip.encode('utf-8'))  # Count actual bytes written

    print(f"Generated file size: {current_size / (1024 * 1024 * 1024):.2f} GB")

# Usage
generate_large_ip_file('large_ip_file.txt', 10)