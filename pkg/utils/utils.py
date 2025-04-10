def read_token(filename):
    try:
        with open(filename, 'r') as file:
            token = file.readline().strip()
        return token, None
    except Exception as e:
        return "", e
