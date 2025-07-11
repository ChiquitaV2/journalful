<template>
  <div class="relative">
    <label v-if="label" :for="id" class="form-label">
      {{ label }}
      <span v-if="required" class="text-error">*</span>
    </label>
    
    <div class="relative">
      <input
        v-if="type !== 'textarea' && type !== 'select'"
        :id="id"
        v-model="inputValue"
        :type="type"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :class="[
          'form-input',
          {
            'border-error focus:border-error focus:ring-red-100': hasError,
            'pr-10': type === 'search'
          }
        ]"
        @blur="$emit('blur', $event)"
        @focus="$emit('focus', $event)"
        @input="handleInput"
      />
      
      <textarea
        v-else-if="type === 'textarea'"
        :id="id"
        v-model="inputValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :rows="rows"
        :class="[
          'form-input resize-vertical',
          {
            'border-error focus:border-error focus:ring-red-100': hasError
          }
        ]"
        @blur="$emit('blur', $event)"
        @focus="$emit('focus', $event)"
        @input="handleInput"
      />
      
      <select
        v-else-if="type === 'select'"
        :id="id"
        v-model="inputValue"
        :disabled="disabled"
        :required="required"
        :class="[
          'form-input',
          {
            'border-error focus:border-error focus:ring-red-100': hasError
          }
        ]"
        @blur="$emit('blur', $event)"
        @focus="$emit('focus', $event)"
        @change="handleInput"
      >
        <option v-if="placeholder" value="" disabled>{{ placeholder }}</option>
        <option
          v-for="option in options"
          :key="option.value"
          :value="option.value"
        >
          {{ option.label }}
        </option>
      </select>
      
      <div v-if="type === 'search'" class="absolute inset-y-0 right-0 pr-3 flex items-center">
        <Icon name="heroicons:magnifying-glass" class="h-5 w-5 text-gray-400" />
      </div>
    </div>
    
    <p v-if="hasError && errorMessage" class="form-error">
      {{ errorMessage }}
    </p>
    
    <p v-else-if="helpText" class="form-help">
      {{ helpText }}
    </p>
  </div>
</template>

<script setup lang="ts">
interface Option {
  value: string | number
  label: string
}

interface Props {
  modelValue?: string | number
  type?: 'text' | 'email' | 'password' | 'search' | 'textarea' | 'select' | 'number' | 'url'
  label?: string
  placeholder?: string
  id?: string
  disabled?: boolean
  required?: boolean
  errorMessage?: string
  helpText?: string
  options?: Option[]
  rows?: number
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  id: () => `input-${Math.random().toString(36).substr(2, 9)}`,
  disabled: false,
  required: false,
  rows: 3
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
  blur: [event: Event]
  focus: [event: Event]
  input: [event: Event]
}>()

const inputValue = computed({
  get: () => props.modelValue || '',
  set: (value) => emit('update:modelValue', value)
})

const hasError = computed(() => !!props.errorMessage)

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
  emit('update:modelValue', target.value)
  emit('input', event)
}
</script>
