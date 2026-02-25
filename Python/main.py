class Cirlce:
    def __init__(self, radius):
        self.radius = radius

    # def area(self):
    #     return 3.14 * self.radius * self.radius
    #
    @property
    def area(self):
        return 3.14 * self.radius * self.radius

    @area.setter
    def area(self, value):
        self.radius = value / 3.14

    def circumference(self):
        return 2 * 3.14 * self.radius

    def __str__(self):
        return f"Circle with radius {self.radius}"


if __name__ == "__main__":
    circle = Cirlce(5)
    # print(circle.area())
    print(circle.area)
    print(circle.circumference())
    circle.area = 10
    print(circle.area)
