<template>
  <div class="space-y-6">
    <!-- 添加里程碑按钮 -->
    <div class="flex justify-end">
      <el-button type="primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        添加里程碑
      </el-button>
    </div>

    <!-- 里程碑列表 -->
    <div v-for="(group, index) in groupedMilestones" :key="index" class="bg-white rounded-lg shadow p-6">
      <h3 class="text-lg font-semibold mb-4 text-indigo-700">{{ group.category }}</h3>
      
      <div class="space-y-4">
        <div v-for="milestone in group.items" 
             :key="milestone.id" 
             class="flex items-center justify-between p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors"
        >
          <div class="flex items-center space-x-4">
            <el-checkbox 
              v-model="milestone.completed"
              @change="handleMilestoneChange(milestone)"
            />
            <div>
              <div class="font-medium">{{ milestone.title }}</div>
              <div class="text-sm text-gray-500">
                {{ milestone.completed ? '完成于: ' + formatDate(milestone.completed_at) : '' }}
              </div>
            </div>
          </div>
          
          <div class="flex items-center space-x-2">
            <el-tag 
              :type="milestone.completed ? 'success' : 'info'"
              size="small"
            >
              {{ milestone.completed ? '已完成' : '进行中' }}
            </el-tag>
            <el-button 
              type="primary" 
              link
              @click="handleEdit(milestone)"
            >
              编辑
            </el-button>
          </div>
        </div>
      </div>

      <!-- 进度统计 -->
      <div class="mt-4 pt-4 border-t border-gray-200">
        <div class="flex justify-between items-center text-sm text-gray-600">
          <span>完成进度</span>
          <span>{{ getCompletionRate(group.items) }}%</span>
        </div>
        <el-progress 
          :percentage="getCompletionRate(group.items)"
          :stroke-width="8"
          :show-text="false"
          status="success"
        />
      </div>
    </div>

    <!-- 添加/编辑里程碑对话框 -->
    <el-dialog
      v-model="showAddDialog"
      :title="editingMilestone ? '编辑里程碑' : '添加里程碑'"
      width="500px"
    >
      <el-form :model="milestoneForm" label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="milestoneForm.title" />
        </el-form-item>
        <el-form-item label="月龄阶段">
          <el-select v-model="milestoneForm.category" placeholder="选择月龄阶段">
            <el-option label="0-3个月" value="0-3个月" />
            <el-option label="4-6个月" value="4-6个月" />
            <el-option label="7-9个月" value="7-9个月" />
            <el-option label="10-12个月" value="10-12个月" />
          </el-select>
        </el-form-item>
        <el-form-item label="完成状态">
          <el-switch v-model="milestoneForm.completed" />
        </el-form-item>
        <el-form-item label="完成时间" v-if="milestoneForm.completed">
          <el-date-picker
            v-model="milestoneForm.completed_at"
            type="datetime"
            placeholder="选择完成时间"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showAddDialog = false">取消</el-button>
          <el-button type="primary" @click="handleSave">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { formatDate } from '../../utils/date'
import { ElMessage } from 'element-plus'

const props = defineProps({
  milestones: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['update', 'create'])

const showAddDialog = ref(false)
const editingMilestone = ref(null)
const milestoneForm = ref({
  title: '',
  category: '',
  completed: false,
  completed_at: null
})

// 按类别分组里程碑
const groupedMilestones = computed(() => {
  const groups = {}
  props.milestones.forEach(milestone => {
    if (!groups[milestone.category]) {
      groups[milestone.category] = {
        category: milestone.category,
        items: []
      }
    }
    groups[milestone.category].items.push(milestone)
  })
  return Object.values(groups)
})

// 计算完成率
const getCompletionRate = (items) => {
  if (!items.length) return 0
  const completed = items.filter(item => item.completed).length
  return Math.round((completed / items.length) * 100)
}

// 处理里程碑状态变化
const handleMilestoneChange = (milestone) => {
  emit('update', {
    ...milestone,
    completed: !milestone.completed,
    completed_at: !milestone.completed ? new Date() : null
  })
}

// 处理编辑
const handleEdit = (milestone) => {
  editingMilestone.value = milestone
  milestoneForm.value = { ...milestone }
  showAddDialog.value = true
}

// 处理保存
const handleSave = () => {
  if (!milestoneForm.value.title || !milestoneForm.value.category) {
    ElMessage.warning('请填写完整信息')
    return
  }

  if (editingMilestone.value) {
    emit('update', {
      ...editingMilestone.value,
      ...milestoneForm.value
    })
  } else {
    emit('create', milestoneForm.value)
  }

  showAddDialog.value = false
  editingMilestone.value = null
  milestoneForm.value = {
    title: '',
    category: '',
    completed: false,
    completed_at: null
  }
}
</script>

<style scoped>
:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #10B981;
  border-color: #10B981;
}

:deep(.el-progress-bar__inner) {
  background-color: #10B981;
}
</style>