import re



def find_mul(data):
    return re.findall(r'mul\(\d{1,3}\,\d{1,3}\)',data.read())
    
def find_digits(string: str):
    a,b = re.findall(r'\d{1,3}', string)
    return int(a),int(b)

data = open('data.data')
results = find_mul(data)

total = 0
print(results)
# for result in results:
#     print(result)