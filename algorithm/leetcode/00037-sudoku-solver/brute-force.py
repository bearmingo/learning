


class Map(object):

    def __init__(self, pic):
        self.pic = pic

    def search_available_nums(self, x, y):
        if self.pic[x][y] != 0:
            return [self.pic[x][y]]

        used = []
        for i in range(0, 9):
            if self.pic[i][y] != 0:
                used.append(self.pic[i][y])
            if self.pic[x][i] != 0:
                used.append(self.pic[x][i])

        x_start = int(x / 3) * 3
        y_start = int(y / 3) * 3

        for i in range(x_start, x_start+3):
            for j in range(y_start, y_start+3):
                if i == x and j == y:
                    continue
                if self.pic[i][j] != 0:
                    used.append(self.pic[i][j])
        
        not_in = []
        for num in range(1, 10):
            if num not in used:
                not_in.append(num)

        return not_in


    def search_next(self, x, y):
        nums = self.search_available_nums(x, y)

        if x == 8 and y == 8 and len(nums) != 0:
            self.pic[x][y] = nums[0]
            return True

        next = x * 9 + y
        while True:
            next += 1
            if next >=81:
                return True

            if self.pic[int(next / 9)][next % 9] == 0:
                break

        for n in nums:
            self.pic[x][y] = n
            if self.search_next(int(next / 9), next % 9):
                return True

        self.pic[x][y] = 0
        return False
    
    def search(self):

        next = 0
        while True:
            if next >=81:
                return True

            if self.pic[int(next / 9)][next % 9] == 0:
                break
            next += 1

        if self.search_next(int(next / 9), next % 9):
            for i in range(0, 9):
                print(self.pic[i])

def main():
    pic = [
        [0,0,8, 0,0,0, 0,0,0],
        [0,0,0, 9,0,0, 2,0,6],
        [7,6,0, 0,1,0, 0,0,0],
        [0,7,0, 0,3,8, 0,0,0],
        [0,0,1, 0,0,0, 9,0,2],
        [0,0,0, 2,0,0, 0,0,0],
        [0,2,3, 0,7,5, 0,1,0],
        [0,0,0, 0,0,0, 0,3,0],
        [4,1,0, 3,9,0, 0,2,0]
    ]

    m = Map(pic)
    m.search()
    print('he')

if __name__ == "__main__":
    main()




        
