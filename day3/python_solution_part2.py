import re
from typing import Match, Iterator

def find_dont(data) -> Iterator[Match[str]]:
    return re.finditer(r"don't\(\)",data)
 

def find_do(data) -> Iterator[Match[str]]:
    return re.finditer(r'do\(\)',data)
    

def find_mul(data) -> Iterator[Match[str]]:
    return re.finditer(r'mul\(\d{1,3}\,\d{1,3}\)',data)
    
def find_digits(string: str) -> tuple[int,int]:
    a,b = re.findall(r'\d{1,3}', string)
    return int(a),int(b)

data = open('data.data').read()

dos  = [span.end(0) for span in find_do(data)]
donts = [span.start(0) for span in find_dont(data)]

valid_indexes = []
valid_indexes.append([0,donts[0]])

dont = 0
for do in dos:
    if do < dont:
        continue
    for dont in donts:
        if dont < do:
            continue
        elif do < dont:
            valid_indexes.append([do, dont])
            break
    ## if last do or dont is a do, takes rest of string and add as valid
    if dont < do:
        valid_indexes.append([do, len(data)-1])

total = 0

for valid_index in valid_indexes:
    start, stop = valid_index[0], valid_index[1]
    results = find_mul(data[start:stop])
    

    for result in results:
        a,b = find_digits(result.group(0))
        total += a*b

print(total)