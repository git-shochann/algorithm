def selection_sort(data):
    length = len(data)
    for index in range(0, length - 1):
        min_index = index

        for pos in range(index + 1, length):
            if data[pos] < data[min_index]:
                min_index = pos
        
        data[index], data[min_index] = data[min_index], data[index]
    
if __name__ == '__main__':
    data = [3, 8, 1, 2, -1, -10, 4, 2]
    selection_sort(data)
    print(data)