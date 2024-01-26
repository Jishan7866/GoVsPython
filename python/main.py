from time import time

def sum(n):
    total=0
    for i in range(1,n+1):
        total+=1
    return total

def main():
    n=1000000000
    st=time()
    total=sum(n)
    en=time()
    print("Time Taken With Python:{:.2f} seconds".format(en-st))
    print("Python Result: ",total)
    
if __name__=="__main__":
    main()