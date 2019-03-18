for i in range(4):
    f = open('input' + str(i) + '.txt', 'w')

    for x in range(1*(i + 1), 11*(i + 1)):
        f.write('{0:2d} {1:3d} {2:4d}\n'.format(x, x**2, x**3) * x)

    f.close()
