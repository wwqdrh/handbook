#include <iostream>

using namespace std;

class Shape
{
    protected:
        int width;
        int height;
    public:
        void setWidth(int w)
        {
            width = w;
        }
        void setheight(int h)
        {
            height = h;
        }
};

class PaintCost
{
    public:
        int getCost(int area)
        {
            return area * 70;
        }
};

class Rectangle: public Shape, public PaintCost
{
    public:
        int getArea()
        {
            return (width * height);
        }
};

int main(void)
{
    Rectangle Rect;
    int area;

    Rect.setWidth(5);
    Rect.setheight(7);

    area = Rect.getArea();

    cout << "Total area: " << Rect.getArea() << endl;

    cout << "Total paint cost: $" << Rect.getCost(area) << endl;

    return 0;
}