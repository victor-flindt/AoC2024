import numpy as np
from numpy import genfromtxt
import pandas as pd
import sys

def is_safe(row) -> bool:
    row = row[~np.isnan(row)]
    if row[0] > row[1]:
        for index in range(0,len(row)-1):
            ## decreasing
            if (row[index] > row[index+1]) & (row[index]-3 <= row[index+1] < row[index]):
                continue
            else:
                return False, f'[Decreasing] {row[index]} > {row[index+1]} & {row[index]-3} < {row[index+1]} < {row[index]}'
        return True, ''
    
    else:
        for index in range(0,len(row)-1):
            ## increasing
            if (row[index] < row[index+1]) & (row[index] < row[index+1] <= row[index]+3):
                continue
            else: 
                return False, f'[increasing] {row[index]} < {row[index+1]} & {row[index]} < {row[index+1]} < {row[index]+3}'
        return True, ''
    


data = pd.read_csv('data.txt', header = None, delimiter=',', names = range(9))
data = data.to_numpy()

# data = np.genfromtxt('data.txt', delimiter = ',')
total_safe = 0


for index, row in enumerate(data):

    safe, msg = is_safe(row)

    if safe:
        print(f'row {index} is Safe')
        total_safe += 1
    else:
        print(f'row {index} is Unsafe, {msg}')


print(f'total: {total_safe}')
    



# print(df)