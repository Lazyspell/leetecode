"""
This module provides solutions to leetcode questions.
"""


def merge(nums1: list[int], m: int, nums2: list[int], n: int) -> list[int]:
    """Merges two sorted list. time nlongn"""
    nums1 = nums1[:m] + nums2
    nums1.sort()
    return nums1


def main():
    """Main function where all the functions will go through"""
    print("This is the main function.")
    print("___________________________________________________________________")
    print("Merge Sorted Array")
    merge_list1 = [1, 2, 3, 0, 0, 0]
    merge_list2 = [2, 5, 6]
    merge_result = merge(merge_list1, 3, merge_list2, 3)
    print(merge_result)
    print("___________________________________________________________________")


if __name__ == "__main__":
    main()
