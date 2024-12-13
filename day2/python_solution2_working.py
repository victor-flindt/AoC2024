import numpy as np
from numpy import genfromtxt
import pandas as pd
import sys

def is_safe_second_check(row) -> bool:
    row = row[~np.isnan(row)]
    if row[0] > row[1]:
        for index in range(0,len(row)-1):
            ## decreasing
            if (row[index] > row[index+1]) & (row[index]-3 <= row[index+1]):
                continue
            else:
                return False
        return True
    
    else:
        for index in range(0,len(row)-1):
            ## increasing
            if (row[index] < row[index+1]) & (row[index+1] <= row[index]+3):
                continue
            else: 
                return False
        return True
    

data = pd.read_csv('data.txt', header = None, delimiter=',', names = range(9))
data = data.to_numpy()

# data = np.genfromtxt('data.txt', delimiter = ',')
total_safe = 0


for index, row in enumerate(data):

    safe, msg = is_safe_second_check(row)

    if safe:
        print(f'row {index} is Safe')
        total_safe += 1
    else:
        for i in range(0,len(row)):
            safe, msg = is_safe_second_check(np.delete(row, index))
            if safe:
                print(f'row {index} is Safe')
            else:
                continue
        print(f'row {index} is Unsafe, {msg}')


print(f'total: {total_safe}')