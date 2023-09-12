/**
 * 【题目】
 * 在一个 106 x 106 的网格中，每个网格上方格的坐标为 (x, y) 。

现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。

每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。

只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。

【分析】
 * 首先，对于1M by 1M的网格内，想要进行任意两点间的寻路几乎是不可能的。本题的关键在于blocked的位置最多只有200个。

正常情况下，200个元素可以围成最大的区域面积是多少呢？肯定是一个中心对称的图形。我们如果从起点走一步，可以发现能走到的范围是一个边长是2、倾斜了45度的正方形。走两步的话，能覆盖的范围是一个边长是3、倾斜了45度的正方形。依次类推，考虑总周长为200的正方形，说明其边长是50，也就是从起点最多走49步。这种正方形最紧凑，是固定周长的条件下能够覆盖的最大的面积（2500）。

于是我们换个角度想：如果从起点出发向外扩散，并且最终可以扩展到超过2500个元素的区域（且没有遇到终点），这说明什么？因为根本没有周长是200的图形可以封闭住这么的面积，所以blocked对于这个起点而言是无效的。也就是说，blocked并不能完全围住起点。

同理，如果我们又发现blocked也不能围住终点的话，那么说明起点和终点势必会相遇。

所以这是一个BFS题，只要从一个点出发开始发散，当visited的网格数目（也就是覆盖的面积）大于2500的时候，就说明这个点并没有被封闭。

有了这个基本思路后，我们需要注意，其实200的周长最大能封闭的面积可以是19900，而不是2500.原因是这200个点可以以45度倾斜地围住一个角。因此0+1+2+...+199 = 19900才是最大的封闭面积。只有发散的区域超过了这个面积，才能保证不被封闭。
*/

using namespace std;

class Solution
{

private:
    using point_t = std::vector<int>;

    static inline uint64_t hash(const point_t &p)
    {
        return p[0] * 1e6 + p[1];
    }

public:
    bool isEscapePossible(vector<vector<int>> &blocked,
                          vector<int> &source,
                          vector<int> &target)
    {
        // blocked points set, uint64_t is selected since it is larger than 1e6*le6
        std::unordered_set<uint64_t> bps;

        for (auto &x : blocked)
            bps.insert(hash(x));

        auto e1 = enclose(source, bps, target);
        auto e2 = enclose(target, bps, source);

        return !e1 && !e2;
    }

private:
    bool enclose(const point_t &s,
                 const std::unordered_set<uint64_t> &b,
                 const point_t &d)
    {
        int max_points = (1 + b.size() - 1) * b.size() / 2;

        std::deque<point_t> Q;
        Q.push_back(s);

        std::unordered_set<uint64_t> visited;
        visited.insert(hash(s));

        while (!Q.empty() && visited.size() <= max_points)
        {
            int &r = Q.front()[0];
            int &c = Q.front()[1];

            static std::vector<std::pair<int, int>> ds{
                {0, 1}, {-1, 0}, {0, -1}, {1, 0}};

            for (auto &dir : ds)
            {
                auto rr = r + dir.first;
                auto cc = c + dir.second;

                std::vector<int> p{rr, cc};

                if (rr < 0 || rr >= 1e6 ||                    // out of bound
                    cc < 0 || cc >= 1e6 ||                    // out of bound
                    visited.find(hash(p)) != visited.end() || // visited
                    b.find(hash(p)) != b.end())               // blocked
                    continue;

                if (p == d) // reached out to destination
                    return false;

                Q.push_back(p);
                visited.insert(hash(p));
            }

            Q.pop_front();
        }
        // no points left in queue, indicating this attempt is enclosed
        return Q.empty();
    }
};

int main() {
    Solution s = Solution();
    s.isEscapePossible();
}