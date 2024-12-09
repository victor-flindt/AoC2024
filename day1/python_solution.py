import pandas as pd
import numpy as np

df = pd.read_csv('data.txt', delimiter=',', names = ['list1','list2'])

list1 = np.sort(df['list1'].to_numpy())
list2 = np.sort(df['list2'].to_numpy())


distance = [abs(cell1-cell2) for cell1,cell2 in zip(list1,list2)]
print(f'answer:{sum(distance)}')