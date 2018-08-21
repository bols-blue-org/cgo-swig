// point/point.h
class point {
	public:
		int x;
		int y;
		void increment() {
			x++;
			y++;
		}
};

void increment(point *p);
