#include <array>
#include <bitset>
#include <iostream>
#include <sstream>
#include <vector>

struct Pt {
  int x = 0;
  int y = 0;

  Pt() : x(0), y(0){};
  Pt(int x_, int y_) : x(x_), y(y_) {}

  std::string string() const {
    std::stringstream ss;
    ss << "(" << x << ", " << y << ")";
    return ss.str();
  }
};

struct Rect {
  Pt bottomLeft;
  Pt topRight;

  int width() const { return topRight.x - bottomLeft.x; }
  int height() const { return topRight.y - bottomLeft.y; }
  int area() const { return width() * height(); }

  std::string string() const {
    std::stringstream ss;
    ss << bottomLeft.string() << "-" << topRight.string();
    return ss.str();
  }
};

const static int MAP_SIZE = 500;

struct Map {
  std::array<std::bitset<MAP_SIZE>, MAP_SIZE> grid;

  void mark(int x, int y) { grid[y][x] = true; }

  void markRect(const Rect &rect) {
    int w = rect.width();
    int h = rect.height();

    for (int y = rect.bottomLeft.y; y < rect.topRight.y; y++) {
      auto &bits = grid[y];
      for (int x = rect.bottomLeft.x; x < rect.topRight.x; x++) {
        bits[x] = true;
      }
    }
  }

  int count() const {
    int total = 0;
    for (const auto &bits : grid) {
      total += bits.count();
    }
    return total;
  }
};

int main() {
  while (true) {
    int n;
    std::cin >> n;
    if (n == 0) {
      break;
    }

    std::vector<Rect> rects(n);
    for (int i = 0; i < n; i++) {
      Pt bottomLeft, topRight;
      std::cin >> bottomLeft.x >> bottomLeft.y >> topRight.x >> topRight.y;
      rects.at(i) = Rect{.bottomLeft = bottomLeft, .topRight = topRight};
    }

    Map map;
    for (const auto &rect : rects) {
      // std::cerr << "rectangle " << rect.string() << ": " << rect.area()
      //           << std::endl;
      map.markRect(rect);
    }

    // std::cerr << "so " << map.count() << std::endl;
    std::cout << map.count() << std::endl;
  }

  return 0;
}
